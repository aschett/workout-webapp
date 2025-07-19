export interface Workout {
    id: number;
    date: string;
    workouts: WorkoutEntry[];
}

export interface WorkoutEntry {
    id: number;
    workoutID: number;
    exerciseID: number;
    exerciseName: string;
    weight: number;
    sets: number;
    reps: number;
}
