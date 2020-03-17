var App = React.createClass({
    componentWillMount: function() {
      this.setState();
    },
    setState: function(){
      var idToken = localStorage.getItem('id_token');
      if(idToken){
        this.loggedIn = true;
      } else {
        localStorage.setItem('id_token', 'bearer secret1')
        this.loggedIn = false;
      }
    },
    render: function() {
      
      if (this.loggedIn) {
        return (<LoggedIn />);
      } else {
        return (<Home />);
      }
    }
  });
  
  var Home = React.createClass({
    authenticate: function(){
      this.loggedIn = true;
      return <Redirect to='${AUTH0_CALLBACK_URL}'/>;
    },
    render: function() {
      return (
      <div className="container">
        <div className="col-xs-12 jumbotron text-center">
          <h1>We Rescue: Animal Rescue</h1>
          <p>Add adopters and adoptees to registry.</p>
          <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
        </div>
      </div>);
    }
  });
  
  var LoggedIn = React.createClass({
    logout : function(){
      localStorage.removeItem('id_token');
      localStorage.removeItem('access_token');
      localStorage.removeItem('profile');
      location.reload();
    },
  
    getInitialState: function() {
      return {
        adopters: [],
        adoptees: []
      }
    },
    componentDidMount: function() {
      this.serverRequest = $.get('http://localhost:3000/adopters', function (result) {
        this.setState({
          adopters: result == "null" ? [] : result,
        });
      }.bind(this));
      this.serverRequest = $.get('http://localhost:3000/adoptees', function (result) {
        this.setState({
          adoptees: result == "null" ? [] : result,
        });
      }.bind(this));

    },
  
    render: function() {
      return (
        <div className="col-lg-12">
          <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
          <h2>Welcome to We Rescue: Animal Rescue</h2>
          <p>Below you'll find the latest adopters and adoptees in the registry. Please complete the form to continue adding to the registry.</p>
          <div className="row">
            
          {this.state.adopters.map(function(adopter, i){
            return <Adopter key={i} adopter={adopter} />
          })}
  
          </div>
          <div className="row">
            {this.state.adoptees.map(function(adoptee, i){
              return <Adoptee key={i} adoptee={adoptee} />
            })}
    
            </div>
          <div className="adopter_form">
            <form action="/adopters" method="post">
              <b>Adopter's Personal Information</b>
              <br/><br/>FIRST NAME
              &nbsp;<input type="text" name="first_name"></input>
              <br/>LAST NAME
              &nbsp;<input type="text" name="last_name"></input>              
              <br/>PHONE
              &nbsp;<input type="text" name="phone"></input>
              <br/>EMAIL
              &nbsp;<input type="text" name="email"></input>
              <br/>GENDER
              &nbsp;<select id="gender" name="gender">
                <option value="female">Female</option>
                <option value="male">Male</option>
                <option value="other">Other</option>
                </select>
              <br/>BIRTHDATE
              &nbsp;<input type="date" name="birthdate"></input>
              <br/>ADDRESS
              &nbsp;<input type="text" name="address"></input>
              <br/>COUNTRY
              &nbsp;<input type="text" name="country"></input>
              <br/>STATE
              &nbsp;<input type="text" name="state"></input>
              <br/>CITY
              &nbsp;<input type="text" name="city"></input>
              <br/>ZIPCODE
              &nbsp;<input type="text" name="zipcode"></input>
              <br/>
              &nbsp;<input type="submit" name="Submit"></input>
            </form>
          </div>
        </div>);
    }
  });
  
  var Adopter = React.createClass({
    waitingToAdopt : function(){
      var adopter = this.props.adopter;
      this.serverRequest = $.post('http://localhost:3000/adopters/' + adopter.id, {status : "waiting"}, function (result) {
        this.setState({adoptionState: "Waiting"})
      }.bind(this));
    },
    Adopted: function(){
      var adopter = this.props.adopter;
      this.serverRequest = $.post('http://localhost:3000/adopters/' + adopter.id, {status : "adopted"}, function (result) {
        this.setState({adoptionState: "Complete"})
      }.bind(this));
    },
    getInitialState: function() {
      return {
        adoptionState: null
      }
    },
    render : function(){
      return(
      <div className="col-xs-4">
        <div className="panel panel-default">
          <div className="panel-heading">{this.props.adopter.FirstName} {this.props.adopter.LastName}<span className="pull-right">{this.state.adoptionState}</span></div>
          <div className="panel-body">
            {this.props.adopter.Phone}
            {this.props.adopter.Email}
          </div>
          <div className="panel-footer">
            <a onClick={this.waitingToAdopt} className="btn btn-default">
              <span className="glyphicon glyphicon-time"></span>
            </a>
            <a onClick={this.Adopted} className="btn btn-default pull-right">
              <span className="glyphicon glyphicon-ok"></span>
            </a>
          </div>
        </div>
      </div>);
    }
  })
  
  ReactDOM.render(<App />,
    document.getElementById('app'));