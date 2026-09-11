CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    owner_team VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE packages (
    id SERIAL PRIMARY KEY,
    ecosystem VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    UNIQUE (ecosystem, name)
);

CREATE TABLE package_versions (
    id SERIAL PRIMARY KEY,
    package_id INT REFERENCES packages(id) ON DELETE CASCADE,
    version_string VARCHAR(100) NOT NULL,
    UNIQUE (package_id, version_string)
);

CREATE TABLE dependency_edges (
    from_version_id INT REFERENCES package_versions(id) ON DELETE CASCADE,
    to_version_id INT REFERENCES package_versions(id) ON DELETE CASCADE,
    PRIMARY KEY (from_version_id, to_version_id)
);

CREATE TABLE project_dependencies (
    project_id INT REFERENCES projects(id) ON DELETE CASCADE,
    package_version_id INT REFERENCES package_versions(id) ON DELETE CASCADE,
    PRIMARY KEY (project_id, package_version_id)
);