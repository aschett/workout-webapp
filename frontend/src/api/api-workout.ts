import type { Workout } from '../types/workout';

export async function getAllWorkouts(): Promise<Workout[]> {
  const res = await fetch('/api/workouts');
  if (!res.ok) throw new Error("Couldn't fech workouts");
  const data = await res.json();
  return data as Workout[];
}


export async function getWorkoutByID(id: number): Promise<Workout> {
  const res = await fetch(`/api/workouts/${id}`);
  if (!res.ok) throw new Error("Workout not available");
  const data = await res.json();
  return data as Workout;
}