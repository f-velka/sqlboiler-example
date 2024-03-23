CREATE TABLE tasks (
    id INT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

INSERT INTO tasks(id, name) VALUES
(1, "task01"),
(2, "task02"),
(3, "task03"),
(4, "task04");

CREATE TABLE task_relations (
    task_id INT NOT NULL,
    depends_on INT NOT NULL,
    PRIMARY KEY (task_id, depends_on),
    FOREIGN KEY (task_id)
        REFERENCES tasks (id),
    FOREIGN KEY (depends_on)
        REFERENCES tasks (id)
);

INSERT INTO task_relations VALUES
(3, 1),
(3, 2),
(4, 3);