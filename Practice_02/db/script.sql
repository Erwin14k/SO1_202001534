-- Resource Table
CREATE TABLE IF NOT EXISTS resource (
	resource INT PRIMARY KEY AUTO_INCREMENT,
    date_resource DATETIME NOT NULL,
    cpu_data FLOAT NOT NULL,
    ram_data FLOAT NOT NULL
);

-- Process Table
CREATE TABLE IF NOT EXISTS process (
	process INT PRIMARY KEY AUTO_INCREMENT,
	pid INT NOT NULL,
    name VARCHAR(60) NOT NULL,
    user INT NOT NULL,
    status VARCHAR(60) NOT NULL,
    ram_percentage FLOAT NOT NULL,
    resource INT NOT NULL,
    parent_process INT NULL,
    FOREIGN KEY (resource) REFERENCES resource(resource) ON DELETE CASCADE,
    FOREIGN KEY (parent_process) REFERENCES process(process) ON DELETE CASCADE
);