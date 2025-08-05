import { ModeToggle } from "@/components/mode-toggle";
import { Link } from "react-router-dom";
import { Toaster } from "sonner";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="max-w-4xl mx-auto p-4">
      <header className="flex justify-between items-center mb-6">
        <Link to="/" className="text-xl font-bold hover:underline">
          Workout App by Andi
        </Link>
        <ModeToggle />
      </header>
      
      <Toaster richColors position="top-center" />

      <main>{children}</main>
    </div>
  );
}