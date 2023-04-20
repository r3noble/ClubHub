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
  - added event database
  - added new structs to act as models for json requests and database models
    - includes: clubAdder.go, events.go, clubs.go


### Front End
- created club search page
  - with redirects to club profiles
- created club profiles 
  - with redirects to club executive member pages
- added ability to join clubs
  - join club buttons only show up when user is logged in
  - ability to join multiple clubs
- added functionality, redirects, permanence to login/logout
- integrated user profile with backend
  - added functionality, redirects to user profile
- fixed calendar to display events with permanence
  -added future ability to integrate events to be handled with our api
- added functionality in user profile page to show clubs that user joined
  - clicking on the club names takes the user to the club profile page
  - shows events for clubs the user is a member of

## Tests

### Back End
- New Unit Tests:
  - addclub_test : tests to ensure clubs are added correctly to database
  - addevent_test : tests to ensure events are added correctly to database
  - getRole_test : tests to ensure that the GetRole API handler works correctly
  - joinClub_test : tests to ensure JoinCLub API handler worrks correctly and that club is formatted successfully in the database
### Front End
- HeaderComponent
  - should create
- AppComponent
  - should have as title 'ClubHub'
  - should create the app
- AuthService
  - should be created
- HomeComponent
  - should create
- CprofileService
  - should be created
- LoginService
  - should be created
- CprofileComponent
  - should create
  - should call cprofileService.getClubInfo with the correct name
  - should set the club stuff when getClubInfo works
  - should log an error when getClubInfo fails
- HttpClient testing
  - SPEC HAS NO EXPECTATIONS works
- HttpClient testing
  - SPEC HAS NO EXPECTATIONS works
- RegisterService
  - should be created
- CalendarComponent
  - should initialize the views property
  - SPEC HAS NO EXPECTATIONS should create
  - should initialize the firstDayOfWeek property
  -  should initialize the shadeUntilCurrentTime property
  - should initialize the dataSource property with specific data
  - should initialize the currentTimeIndicator property
  - should initialize the view property
  - should initialize the dataSource property with data
- RegisterService
  - should be created
  - RegisterComponent
  - should navigate to home page on cancel
  - should create
- HttpClient testing
  - SPEC HAS NO EXPECTATIONS works
- ClubComponent
  - should create
 - CalComponent
  - should create
- PublicprofileComponent
  - should create
- StepperComponent
  - should create
- HttpClient testing
  - SPEC HAS NO EXPECTATIONS works
- LoginComponent
  - should show error message on incorrect creds
  - should navigate to register page
  - should login
  - should create
- ProfileService
  - should be created
- ProfileComponent
  - should set the clubs property to "No clubs joined yet!" if no clubs joined
  - should create
  - should set the name and email properties from the user object
  - should set the clubs property from the user object if it exists
- FooterComponent
  - should create
- PublicprofileService
  - should be created

## Updated Back End Documentation
- ## GET
 #### able to search and find existing users
- func (a *App) GetUserByID(id string): able to grab user profile by both id, name, and email
- func (a *App) IdHandler: Help function to the previous function
- getClubHandler: Grab club information and pass back based on club name
- profileHandler: handles user list within clubs and lists 

#### able to return detailed issues within searching
- ## POST
#### adding user to DB
 func (a *App) AddUserHandler <- gets passed JSON information
#### allow to add users to clubs
- joinClubHandler: checks clubs csv within user struct and adds clubs if possible
- getRoleHandler: checks an existing users membership within a club and returns their membership
- 
#### login credentials posting
func (a *App) loginHandler <- passed username and password creds with JSON
- ## TESTING
#### func HealthCheck returns plain text if API is running.
