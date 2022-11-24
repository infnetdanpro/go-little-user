CREATE TABLE users_roles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    role_id INTEGER,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_role_idd FOREIGN KEY(role_id) REFERENCES roles(id)
);