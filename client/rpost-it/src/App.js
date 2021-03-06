import React, { Component } from 'react';
import './App.css';
import { HashRouter as Router, Route, Switch,  } from 'react-router-dom';
import Home from './Home';

import About from './About';
import NoMatch from './NoMatch';
import Layout from './components/Layout';
import NavigationBar from './components/NavigationBar';


import Signup from './signup';
class App extends Component {
    render() {
        return (
            <React.Fragment>
                <NavigationBar />
                <Layout>
                    
                    <Router>
                        <Switch>
                            <Route exact path="/" component={Home} />
                            <Route path="/login" component={About} />
                            <Route path="/signup" component={Signup} />
                            <Route component={NoMatch} />
                        </Switch>
                    </Router>
                </Layout>
            </React.Fragment>
        );
    }
}

export default App;
