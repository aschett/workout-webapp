CREATE TABLE exercises (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE workouts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    date DATE NOT NULL UNIQUE
);

CREATE TABLE workoutEntries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    workoutID INT NOT NULL,
    exerciseID INT NOT NULL,
    weight FLOAT NOT NULL,
    sets INT NOT NULL,
    reps INT NOT NULL,
    FOREIGN KEY (workoutID) REFERENCES workouts(id) ON DELETE CASCADE,
    FOREIGN KEY (exerciseID) REFERENCES exercises(id) ON DELETE CASCADE
);

INSERT INTO exercises (name) VALUES
('Bench Press'),
('Squat'),
('Deadlift'),
('Overhead Press'),
('Barbell Row'),
('Pull Up'),
('Dumbbell Curl'),
('Tricep Pushdown'),
('Lat Pulldown'),
('Leg Press');

INSERT INTO workouts (date) VALUES
('2024-06-20'),
('2024-06-22'),
('2024-06-24'),
('2024-06-26'),
('2024-06-28'),
('2024-06-30'),
('2024-07-02'),
('2024-07-04'),
('2024-07-06'),
('2024-07-08');

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(1, 1, 100, 3, 10),
(1, 2, 120, 3, 8),
(1, 5, 90, 3, 10);

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(2, 4, 60, 3, 12),
(2, 7, 20, 3, 15),
(2, 8, 35, 3, 15);

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(3, 3, 150, 2, 5),
(3, 5, 95, 3, 10),
(3, 6, 0, 3, 10);

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(4, 2, 130, 3, 8),
(4, 1, 105, 3, 10),
(4, 9, 50, 3, 12);

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(5, 10, 180, 4, 10),
(5, 3, 155, 2, 5),
(5, 5, 100, 3, 10);

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(6, 4, 65, 3, 12),
(6, 7, 22.5, 3, 15),
(6, 8, 37.5, 3, 15);

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(7, 1, 110, 3, 10),
(7, 2, 135, 3, 8),
(7, 5, 105, 3, 10);

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(8, 3, 160, 2, 5),
(8, 6, 0, 3, 10),
(8, 9, 55, 3, 12);

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(9, 10, 190, 4, 10),
(9, 1, 115, 3, 10),
(9, 7, 25, 3, 15);

INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps) VALUES
(10, 2, 140, 3, 8),
(10, 4, 70, 3, 12),
(10, 8, 40, 3, 15);
