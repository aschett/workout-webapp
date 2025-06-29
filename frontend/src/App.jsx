import { BrowserRouter, Routes, Route } from 'react-router-dom';
import WorkoutList from './components/WorkoutList';
import WorkoutPage from './components/WorkoutPage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<WorkoutList />} />
        <Route path="/workouts/:id" element={<WorkoutPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
