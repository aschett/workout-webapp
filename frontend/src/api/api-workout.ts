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

export async function createWorkout(workout: { date: string }) {
  const body = new URLSearchParams();
  body.append('date', workout.date);

  const res = await fetch('/api/workouts', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: body.toString(),
  });

  if (res.status === 409) {
    throw new Error('Workout with date '+ workout.date +' already exists!')
  }

  if (!res.ok) throw new Error('Failed to create workout');
}
