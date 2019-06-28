# Waiting mysql
until mysqladmin ping -uroot -ppassword -h mysql_user --silent; do
    echo 'waiting for mysqld to be connectable...' && sleep 3;
done

mysqladmin -uroot -ppassword -h mysql_user create bh_user

goose -env=compose -path=. up

mysql -f  -h mysql_user -uroot -ppassword bh_user < testdata/testdata.sql