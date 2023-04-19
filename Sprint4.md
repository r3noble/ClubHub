# Sprint 4

## Work Completed

### Back End
- added club database
- added functions to
  -  get user by name and email functional
  -  add clubs
  -  view clubs
  -  get club profiles
  -  add club to user 


### Front End
- created club search page
  - with redirects to club profiles
- created club profiles 
  - with redirects to club executive member pages
- added functionality, redirects, permanence to login/logout
- integrated user profile with backend
  - added functionality, redirects to user profile
- fixed calendar to display events with permanence
  -added future ability to integrate events to be handled with our api


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
