import React, { Component } from 'react';
import { Dropdown,NavItem } from 'react-bootstrap';
import styled from 'styled-components';
const Styles = styled.div`
    
`;
export default class LoginState extends Component {
    constructor(props) {
        super();
        this.props = props;
    }
    render() {
        return (
            <Styles>
                <Dropdown>
                    <Dropdown.Toggle variant="info" id="dropdown-basic">
                        {this.props.loginName}
                    </Dropdown.Toggle>

                    <Dropdown.Menu>
                        <Dropdown.Item href="#/action-1">Accounts</Dropdown.Item>
                        <Dropdown.Item href="#/action-2">Another action</Dropdown.Item>
                        <Dropdown.Item href="#/action-3">Something else</Dropdown.Item>
                    </Dropdown.Menu>
                </Dropdown>
            </Styles>
        );
    }
}
