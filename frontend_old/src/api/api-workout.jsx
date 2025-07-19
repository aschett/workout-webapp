export async function getAllWorkouts() {
  const res = await fetch('/api/workouts');
  return res.json();
}

export async function getWorkoutByID(id) {
  const res = await fetch(`/api/workouts/${id}`);
  return res.json();
}