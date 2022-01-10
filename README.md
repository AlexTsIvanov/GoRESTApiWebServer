# GoLang REST API Web Server

Order Management Web Application for restaurants

## Reason Behind the Project

Dark(Ghost) kitchens are innovative restaurants that operate for delivery only. That way restaurant operating costs are greatly reduced by removing the need for waiters, suitable location etc. The main method of getting orders is through their online presence. 

## Project Description

This project provides a way for customers to get familiar with the menu, place orders, pay online and track their orders.
The main user roles are:
•	Anonymous User – can only view the menu pages. 
•	Customer (type of Registered User) – can choose items to order and place orders.
•	Kitchen Staff (type of Registered User) – can view pending orders, can work on an order, can complete an order by giving it to delivery driver.
•	Administrator (type of Registered User) – can manage (create, edit user data and delete) all Registered Users, as well as manage menu listings.

## Development Process

• Created a MySQL database using defined models using GORM.
• Created router using the Gorilla framework.
• Added endpoints to provide for different interactions with the database.
• Created functions to allow for new customers to be able to create an account or log into existing one.
• JWT is generated when user signs in where the user role is encoded.
• Added an authorization middleware to check if client is allowed to access a certain endpoint.

## Future Work

Finish adding all the needed endpoint permissions.
Create front-end which consumes the data.
