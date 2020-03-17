var App = React.createClass({
    componentWillMount: function() {
      this.setupAjax();
      this.parseHash();
      this.setState();
    },
    setupAjax: function() {
      $.ajaxSetup({
        'beforeSend': function(xhr) {
          if (localStorage.getItem('access_token')) {
            xhr.setRequestHeader('Authorization',
                  'Bearer ' + localStorage.getItem('access_token'));
          }
        }
      });
    },
    parseHash: function(){
      this.auth0 = new auth0.WebAuth({
        domain:       AUTH0_DOMAIN,
        clientID:     AUTH0_CLIENT_ID
      });
      this.auth0.parseHash(window.location.hash, function(err, authResult) {
        if (err) {
          return console.log(err);
        }
        if(authResult !== null && authResult.accessToken !== null && authResult.idToken !== null){
          localStorage.setItem('access_token', authResult.accessToken);
          localStorage.setItem('id_token', authResult.idToken);
          localStorage.setItem('profile', JSON.stringify(authResult.idTokenPayload));
          window.location = window.location.href.substr(0, window.location.href.indexOf('#'))
        }
      });
    },
    setState: function(){
      var idToken = localStorage.getItem('id_token');
      if(idToken){
        this.loggedIn = true;
      } else {
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
      this.webAuth = new auth0.WebAuth({
        domain:       AUTH0_DOMAIN,
        clientID:     AUTH0_CLIENT_ID,
        scope:        'openid profile',
        audience:     AUTH0_API_AUDIENCE,
        responseType: 'token id_token',
        redirectUri : AUTH0_CALLBACK_URL
      });
      this.webAuth.authorize();
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
        adoptees: [],
      }
    },
    componentDidMount: function() {
      this.serverRequest = $.get('http://localhost:3000/adopters', function (result) {
        this.setState({
          adopters: result,
        });
      }.bind(this));
      this.serverRequest = $.get('http://localhost:3000/adoptees', function (result) {
        this.setState({
          adoptees: result,
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