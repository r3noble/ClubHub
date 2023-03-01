# Sprint 2 Overview
## Work Completed in Sprint 2
- Integrated front end and back end to perform a simple (login or profile component)
### Front End
- Added a calendar to the calendar page that we struggled for a long time with in Sprint 1.
- Added a profile page for a user. (club to come later)
- Implemented a registration page that routes from the login page. (no button in the navbar in header file so had to route through login) Only takes in valid inputs from the user such as only allowing UF emails.
- Added a simple cypress test.
- Added unit tests.

### Back End
- Added user searching functionality which looks for existing users based on primary name tag
- Linked a profile search to searched for existing users
- Added health checks to APIs to test running.
- Added unit tests.
# Testing

## Front End Unit Tests


## Front End Cypress Test
Wrote Cypress test that involved an up-down counter with buttons, implemented the following checks to accompany:
1. Make sure counter initializes to zero every time
2. Makes sure it is possible to change the counter value to a specific number
3. Makes sure that the '-' button decrease the counter, and '+' button increases the counter
4. Makes sure that clicking makes a change event happen with the counter

## Back End Unit Tests
1. Make sure we can grab a user by name(ID)
2. Testing login handler with known login information
3. Adding user to backend database
# Back End Documentation

