import React, { Component } from 'react'
import CardComponent from './components/Card';

class Home extends Component {
    render() {
        return (
            <div>
                <CardComponent/>
                <CardComponent/>
                <CardComponent/>
                <CardComponent/>
                <CardComponent/>
            </div>
        )
    }
}

export default Home