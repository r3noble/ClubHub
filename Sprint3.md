# Sprint 3

## Work Completed
#### Front End
- Integrated a post with the backend to improve our front/back communications
- Performed a lot of debugging with our testing issues that arose in the last sprint
- Added new tests
- Now have about 30 working tests, compared to none on last sprint
- Added register functionality with back end
- Added login functionality with back end

#### Back End
- Reworked our database to make it functional
- Cleaned up the repository (more reworking)
- 
## Unit Tests

#### Back End
- loginhandler_test.go
    - tests our login API to ensure it is receiving and interpretting json request correctly and accessing correct user from database correctly
- adduser_test.go
    - tests our addUser API to ensure it is receiving/interpretting json correctly and then properly adding user to database

#### Front End

##### ProfileService
- should be created

##### LoginComponent
- should call loginService login method on submit
- should navigate to register page on register button click
- should set error message and navigate to profile page on failed login
- should create

##### RegisterComponent
- should navigate to home page on cancel
- should create

##### LoginService
- should be created

##### CalendarComponent
- should initialize the shadeUntilCurrentTime property
- should initialize the dataSource property with data
- should create
- should initialize the firstDayOfWeek property
- should initialize the currentTimeIndicator property
- should initialize the views property
- should initialize the view property
- should initialize the dataSource property with specific data

##### ProfileComponent
- should create

##### FooterComponent
- should create

##### StepperComponent
- should create
- 
##### HomeComponent
- should create

##### HttpClientTesting
- works

##### AppComponent
- should create the app
- should have title 'ClubHub'

##### HeaderComponent
- should create


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
