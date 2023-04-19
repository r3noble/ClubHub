# Sprint 4

## Work Completed

### Back End
- added club database
- added functions to
  -  add clubs
  -  view clubs/ create club profiles
  -  add club to user 

### Front End
- created club search page with redirects to club profiles
- created club profiles 
- added functionality, redirects, permanence to login
- integrated user profile with backend
- added functionality, redirects to user profile
- fixed calendar to display events with permanence


## Tests

### Back End

### Front End

## Updated Back End Documentation
- ## GET
 #### able to search and find existing users
func (a *App) GetUserByID(id string)
func (a *App) IdHandler
#### able to return detailed issues within searching
- ## POST
#### adding user to DB
 func (a *App) AddUserHandler <- gets passed JSON information
#### login credentials posting
func (a *App) loginHandler <- passed username and password creds with JSON
- ## TESTING
#### func HealthCheck returns plain text if API is running.
