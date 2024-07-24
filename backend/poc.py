from argparse import ArgumentParser
import struct
import matplotlib.pyplot as plt
import numpy as np
import math
import cv2
from sklearn.cluster import KMeans
import skimage
from skimage import measure


def read_elevation_file(file_path, width):
    # Read the binary file
    with open(file_path, 'rb') as file:
        # Read all the data as 32-bit floats
        data = np.fromfile(file, dtype=np.float32)
    
    # Calculate the height of the image
    height = len(data) // width
    
    # Reshape the data into a 2D array (image)
    elevation_data = data.reshape((height, width))
    
    return elevation_data

def compute_derivative(elevation_data):
    # Compute the derivative of the elevation data
    dx, dy = np.gradient(elevation_data)
    
    return dx, dy

def derivative_to_normal_map(dx, dy):
    # Compute the normal map from the derivatives
    normal_map = np.dstack((-dx, -dy, np.ones_like(dx)))
    normal_map /= np.linalg.norm(normal_map, axis=2)[:, :, None]
    
    return normal_map

def normal_to_uv_color_map(normal_map):
    # Compute the UV color map from the normal map
    u = 0.5 + np.arctan2(normal_map[:, :, 1], normal_map[:, :, 0]) / (2 * np.pi)
    v = 0.5 - np.arcsin(normal_map[:, :, 2]) / np.pi
    return u, v

def cluster(u,v):
    kmeans = KMeans(n_clusters=5, random_state=0, algorithm="auto").fit(np.dstack((u, v)).reshape(-1, 2))
    labels = kmeans.labels_
    labels_im = labels.reshape(u.shape)  
    return labels_im

def uv_to_angle_normalized(u, v):
    u_norm = (u - u.min()) / (u.max() - u.min())
    v_norm = (v - v.min()) / (v.max() - v.min())
    dir = np.dstack((u_norm, v_norm, np.zeros_like(u)))
    return dir

if __name__ == '__main__':
    parser = ArgumentParser()
    parser.add_argument('--elevation', type=str, required=True)
    parser.add_argument('--image', type=str, required=True)
    args = parser.parse_args()

    image = plt.imread(args.image)
    (width, height) = image.shape[:2]

    elevation_data = read_elevation_file(args.elevation, width)

    # num_labels, labels_im = detect_roof_segments(elevation_data)
    # plot_segments(elevation_data, labels_im, image)

    dx, dy = compute_derivative(elevation_data)
    normal_map = derivative_to_normal_map(dx, dy)
    u, v = normal_to_uv_color_map(normal_map)
    labels_im = cluster(u,v)

    dir = uv_to_angle_normalized(u, v)
    # mask when dir is 0.6
    dir_masked = np.where(np.abs(dir[:,:,0] - 0.5) < 0.1, 1, 0)

    connected_components = measure.label(dir_masked, background=0)
    biggest_component = np.argmax(np.bincount(connected_components.flat)[1:]) + 1
    labels_masked = np.where(connected_components == biggest_component, 1, 0)

    labels_masked_smooth = cv2.GaussianBlur(labels_masked.astype(np.float32), (11, 11), 0)

    plt.figure()
    plt.imshow(image)
    plt.imshow(labels_masked_smooth, alpha=0.5)
    plt.colorbar()
    plt.show()

    