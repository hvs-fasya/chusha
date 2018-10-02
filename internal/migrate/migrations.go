package migrate

import "github.com/rubenv/sql-migrate"

func getSource() (migrations *migrate.MemoryMigrationSource) {
	migrations = &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id: "1",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS users(
									id bigserial not null,
									email text,
									phone text,
									nickname text,
									name text,
									lastname text,
									type text not null,
									pswd_hash text,
									primary key(id)
								);
					INSERT INTO users (nickname, name, type) VALUES ('Nina', 'Nina', 'admin');
					INSERT INTO users (nickname, name, type) VALUES ('admin', 'Lena', 'admin');`,
				},
				Down: []string{"DROP TABLE users"},
			},
			&migrate.Migration{
				Id: "2",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS posts(
									id bigserial not null,
									title text not null,
									content text,
									published_at timestamp not null,
									deleted_at timestamp,
									primary key(id),
									unique (title)
								)`,
				},
				Down: []string{"DROP TABLE posts"},
			},
			&migrate.Migration{
				Id: "3",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS comments(
									id bigserial not null,
									post_id int not null,
									comment_id int,
									user_id int not null,
									content text,
									hidden bool default false,
									primary key(id),
									CONSTRAINT comments_post_id_fkey foreign key (post_id) REFERENCES posts(id) ON DELETE CASCADE,
									CONSTRAINT comments_comment_id_fkey foreign key (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
									CONSTRAINT comments_user_id_fkey foreign key (user_id) REFERENCES users(id) ON DELETE CASCADE
								)`,
				},
				Down: []string{"ALTER TABLE comments DROP CONSTRAINT comments_user_id_fkey; " +
					"ALTER TABLE comments DROP CONSTRAINT comments_comment_id_fkey; " +
					"ALTER TABLE comments DROP CONSTRAINT comments_post_id_fkey; " +
					"DROP TABLE comments;"},
			},
			&migrate.Migration{
				Id: "4",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS tab_types(
									id bigserial not null,
									type text not null,
									primary key(id)
						);
					INSERT INTO tab_types (type) VALUES ('blog');`,
				},
				Down: []string{"DROP TABLE tab_types;"},
			},
			&migrate.Migration{
				Id: "5",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS tabs(
									id bigserial not null,
									title text not null,
									user_type_visible text[],
									tab_type_id int,
									primary key(id),
									CONSTRAINT tabs_types_tab_type_id_fkey foreign key (tab_type_id) REFERENCES tab_types(id) ON DELETE CASCADE
								);
					INSERT INTO tabs (title, user_type_visible, tab_type_id) VALUES ('ЗАПИСИ', '{"all"}', 
						(SELECT id FROM tab_types WHERE type='blog')
					);`,
				},
				Down: []string{"ALTER TABLE tabs DROP CONSTRAINT tabs_types_tab_type_id_fkey; DROP TABLE tabs;"},
			},
			//			&migrate.Migration{
			//				Id: "6",
			//				Up: []string{
			//					`CREATE INDEX IF NOT EXISTS "reports_time_start_btree_idx" ON reports USING btree (time_start);
			//					CREATE INDEX IF NOT EXISTS "reports_time_end_btree_idx" ON reports USING btree (time_end)`,
			//				},
			//				Down: []string{"DROP INDEX IF EXISTS reports_time_end_btree_idx; DROP INDEX IF EXISTS reports_time_start_btree_idx;"},
			//			},
		},
	}
	return
}
