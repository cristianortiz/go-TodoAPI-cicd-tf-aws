db.createCollection("users");
db.users.createIndex({ email: 1 }, { unique: true });

db.createCollection("projects");
db.projects.createIndex({ name: 1 }, { unique: true });

db.createCollection("tasks");
db.tasks.createIndex({ projectID: 1, assignedTo: 1 });

db.createCollection("roles");
db.roles.createIndex({ name: 1 }, { unique: true });
