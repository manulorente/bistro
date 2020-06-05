const AUTH0_CLIENT_ID = "GR6mGA4jOvDZINkW6yIarBuGlSGU1rVB";
const AUTH0_DOMAIN = "dev-ju4b9g5i.eu.auth0.com";
const AUTH0_CALLBACK_URL = location.href;
const AUTH0_API_AUDIENCE = "https://bistro/api";

class App extends React.Component {
    parseHash() {
        this.auth0 = new auth0.WebAuth({
          domain: AUTH0_DOMAIN,
          clientID: AUTH0_CLIENT_ID
        });
        this.auth0.parseHash(window.location.hash, (err, authResult) => {
          if (err) {
            return console.log(err);
          }
          if (
            authResult !== null &&
            authResult.accessToken !== null &&
            authResult.idToken !== null
          ) {
            localStorage.setItem("access_token", authResult.accessToken);
            localStorage.setItem("id_token", authResult.idToken);
            localStorage.setItem(
              "profile",
              JSON.stringify(authResult.idTokenPayload)
            );
            window.location = window.location.href.substr(
              0,
              window.location.href.indexOf("#")
            );
          }
        });
      }
        
      setup() {
        $.ajaxSetup({
          beforeSend: (r) => {
            if (localStorage.getItem("access_token")) {
              r.setRequestHeader(
                "Authorization",
                "Bearer " + localStorage.getItem("access_token")
              );
            }
          }
        });
      }
      setState() {
        let idToken = localStorage.getItem("id_token");
        if (idToken) {
          this.loggedIn = true;
        } else {
          this.loggedIn = false;
        }
      }
        
      componentWillMount() {
        this.setup();
        this.parseHash();
        this.setState();
      }    
    render() {
      if (this.loggedIn) {
        return (<LoggedIn />);
      } else {
        return (<Home />);
      }
    }
  }

class Home extends React.Component {
    constructor(props) {
        super(props);
        this.authenticate = this.authenticate.bind(this);
      }
      authenticate() {
        this.WebAuth = new auth0.WebAuth({
          domain: AUTH0_DOMAIN,
          clientID: AUTH0_CLIENT_ID,
          scope: "openid profile",
          audience: AUTH0_API_AUDIENCE,
          responseType: "token id_token",
          redirectUri: AUTH0_CALLBACK_URL
        });
        this.WebAuth.authorize();
      }
    render() {
        return (
        <div className="container">
            <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1>Bistro</h1>
            <p>Gestión del menu</p>
            <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Acceder</a>
            </div>
        </div>
        )
    }
}

class LoggedIn extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        products: []
      }
      this.serverRequest = this.serverRequest.bind(this);
      this.logout = this.logout.bind(this);
    }

    logout() {
        localStorage.removeItem("id_token");
        localStorage.removeItem("access_token");
        localStorage.removeItem("profile");
        location.reload();
    }
    
    serverRequest() {
    $.get("http://localhost:3000/api/products/", res => {
        this.setState({
        jokes: res
        });
    });
    }
    
    componentDidMount() {
    this.serverRequest();
    } 

    render() {
      return (
        <div className="container">
          <div className="col-lg-12">
            <br />
            <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
            <h2>Bistro</h2>
            <p>Menu virtual</p>
            <div className="row">
              {this.state.products.map(function(products, i){
                return (<ID key={i} name={name} />);
              })}
            </div>
          </div>
        </div>
      )
    }
}

class Product extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        available: "Yes",
        products: []
      }
      this.name = this.name.bind(this);
    }
      
    name() {
        let product = this.props.product;
        this.serverRequest(product);
    }

    serverRequest(joke) {
        $.post(
          "http://localhost:3000/api/products/",
          res => {
            console.log("res... ", res);
            this.setState({ available: "Yes", products: res });
            this.props.products = res;
          }
        );
      }
      
    render() {
      return (
        <div className="col-xs-4">
          <div className="panel panel-default">
            <div className="panel-heading">#{this.props.product.id} <span className="pull-right">{this.state.available}</span></div>
            <div className="panel-body">
              {this.props.product.cat}
            </div>
            <div className="panel-footer">
              {this.props.product.name} Likes &nbsp;
              <a onClick={this.name} className="btn btn-default">
                <span className="glyphicon glyphicon-thumbs-up"></span>
              </a>
            </div>
          </div>
        </div>
      )
    }
  }

 ReactDOM.render(<App />, document.getElementById('app'));
