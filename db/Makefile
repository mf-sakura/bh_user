prepare-db:
	mysqladmin -u root create bh_user
	mysqladmin -u root create bh_user_test

down:
	goose -path=. down
	goose -path=. -env=test down

up:
	goose -path=. up
	goose -path=. -env=test up