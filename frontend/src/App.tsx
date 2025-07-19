import { BrowserRouter, Routes, Route } from "react-router-dom";
import ListAllWorkouts from "./components/ListAllWorkouts";
import ListExercisesForWorkout from "./components/ListExercisesForWorkout";
import Layout from "./components/Layout";
import { ThemeProvider } from "@/components/theme-provider";

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="/" element={<ListAllWorkouts />} />
            <Route path="/workouts/:id" element={<ListExercisesForWorkout />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
