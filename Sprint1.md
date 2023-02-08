User stories
- A Club Executive wants to..<br>
 create a profile<br>
 advertise events<br>
  =>manage their members<br>
  =>allow members to be blocked from viewing club page<br>
  =>allow inactive members to be removed from followers list<br>
 add a bio and other information<br>
  =>create bio header underneath club name/logo on club page<br>
 be able to remind members of an upcoming meeting.<br>
  =>The user is able to send out an email to the account emails members used to register<br>
 be able to promote and demote other members into positions in the club<br>
  =>some executive members can have privileges to plan events and edit the club description<br> 
  
- A school admin wants to...<br>
 limit who can view club events<br>
  =>only students with ufl.edu domains are able to view or join clubs on ClubHub 
 mointor club activities<br>
  =>allow access to view all club pages<br>

- A club member wants to...<br>
 give feedback<br>
  =>add portal in club page for public comments (maybe like Facebook?)<br>
 view events<br>
  =>create timeline where events will be listed<br>
  =>able to click on events to bring up a page on the event<br>
  =>allow for likes & comments to be public under the events on the timeline<br>
 
- A prospective member wants to...<br>
 enter their schedule<br>
  =>give users access to a personal calender that events can be added to<br>
 see how many active members are in the club<br>
  =>add a "followers" tab (like Instagram)<br>
 see links to the clubs' recent social media<br>
  =>add link availability to bio (like Instagram)<br>
 
- A company wants to...<br>
 see a resume portal from the club<br>
  =>allow club members to upload resumes to a pool that only companies can view<br>
  =>make pool available to view when company user views club web page<br>


What issues your team planned to address?

-The front-end team planned on addressing a home page and the login page, with working tabs to navigate between the pages and a login that allowed the user to enter their credentials. The team also discussed a profile where the user would be able to see what clubs and events they have expressed interest in, as well as a calendar page.

-The back-end team planned to address the creation of a server to support the web application, the creation of a sign-up and sign-in functionality that would allow users to create an account with a certain type (member, admin, corporate). To accompany this functionality, we also plan to implement a database that would store the passwords cryptically, as well as creating objects for the various user types and a club object as well.


Which ones were successfully completed?

-The front-end team successfully completed a home page, with a couple working tabs (through routing), and the beginning of a login page (not connected with the back end).

-The back-end team successfully created a server which will support a basic web application, however updates will be required in the future to support more complex functions. A sign-in and sign-up function were also successfully created, however as we will later discuss they may not be able to be used at this time due to unsuccessful implementation of the database. Basic user and club objects were also made to support very basic user and club functions such as storing and updating names, types, etc.


Which ones didn't and why?

-The front-end team did not complete the profile page or the calendar page, as learning to implement routing took much longer than anticipated and debugging the home and login page used up the time for this sprint.

-The back-end team was unable to implement the aforementioned database system. This was due to a fundamental lack of understanding on how PostgreSQL works which required much more troubleshooting time than anticipated. We plan on implementing GORM to handle this in the future to allow for easier and more basic accessing.
