# Sprint 3

## Work Completed
#### Front End
- Integrated a post with the backend to improve our front/back communications
- Performed a lot of debugging with our testing issues that arose in the last sprint
#### Back End
- Reworked our database to make it functional
- Cleaned up the repository (more reworking)
- 
## Unit Tests
#### Front End
- 
#### Back End
- 
## Updated Backend API Documentation
- ## GET
 ### able to search and find existing users
func (a *App) GetUserByID(id string)
func (a *App) IdHandler
### able to return detailed issues within searching
- ## POST
### adding user to DB
 func (a *App) AddUserHandler <- gets passed JSON information
### login credentials posting
func (a *App) loginHandler <- passed username and password creds with JSON
- ## TESTING
### func HealthCheck returns plain text if API is running.
