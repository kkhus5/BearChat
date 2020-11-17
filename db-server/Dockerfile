# What image are we pulling? What version do we want?
# FROM <YOUR CODE HERE>
FROM mysql/mysql-server:latest

# copy something from our directory to our image, like we have shell access to our container's directory
COPY initdb.sql /docker-entrypoint-initdb.d/

# What is the default MySQL port?
# EXPOSE <YOUR CODE HERE> 
EXPOSE 3306