import { BrowserRouter, Routes, Route } from "react-router-dom";
import ListAllWorkouts from "./components/ListAllWorkouts";
import Layout from "./components/Layout";
import { ThemeProvider } from "@/components/theme-provider";

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="/" element={<ListAllWorkouts />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
