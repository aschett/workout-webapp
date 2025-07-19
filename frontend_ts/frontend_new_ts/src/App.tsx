import { BrowserRouter, Routes, Route } from "react-router-dom";

import Layout from "./components/Layout";
import { ThemeProvider } from "@/components/theme-provider";

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <BrowserRouter>
        <Layout>
          Hi
        </Layout>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
