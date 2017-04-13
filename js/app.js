import React from 'react'
import ReactDOM from 'react-dom'

import {
  BrowserRouter as Router,
  StaticRouter, // for server rendering
  Route,
  Link
} from 'react-router-dom'

import DepositsList from './components/deposits_list'

class App extends React.Component {
  render() {
    console.log("App render()")

    return <div>
        <div className="container">
        </div>
      </div>
  }
}

class Routes extends React.Component {
  render() {
    console.log('Routes');

    return <Router>
    <div>
        <ul>
          <li><Link to="/">Home</Link></li>
          <li><Link to="/deposits">Deposits</Link></li>
        </ul>

        <hr/>

        <Route exact path="/" component={App}/>
        <Route path="/deposits" component={DepositsList}/>
      </div>
    </Router>
  }
}

ReactDOM.render(<Routes/>, document.getElementById('app'));