package db

const dbSchema = `
	create table users (
		user_id			int primary key generated always as identity,
		passport_number	text not null unique,
		user_name 		text not null,
		surname 		text not null,
		patronymic		text,
		address			text not null,
		deleted			bool not null default false 
	);

	create table tasks (
		task_id			int primary key generated always as identity,
		executor		int references users(user_id),
		task_name		text not null,
		created_at		timestamp not null,
		completed_at	timestamp
	);
`
