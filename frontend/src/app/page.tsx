import NewAnalysisForm from "./components/NewAnalysisForm";
import StatusIndicator from "./components/StatusIndicator";


export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <NewAnalysisForm />
      <StatusIndicator />
    </main>
  );
}
