import React, { Component } from 'react';
import { Nav, Navbar, Form, FormControl, Button, Col, InputGroup } from 'react-bootstrap';
import styled from 'styled-components';

const Styles = styled.div`
    .navbar {
        background-color: #222;
    }
    .navbar-brand,
    .navbar-nav .nav-link {
        color: #bbb;
        &:hover {
            color: white;
        }
    }
    .navbar-collapse {
        padding: 5px;
    }
`;

export default class NavigationBar extends Component {
    render() {
        return (
            <Styles>
                <Navbar expand="lg">
                    <Navbar.Brand className="mb-2" href="/">
                        rpost-it
                    </Navbar.Brand>
                    <Navbar.Toggle area-aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Form>
                            <Form.Row className="align-items-center">
                                <Col className="ml-auto">
                                    <Form.Label htmlFor="inlineFormInputName" srOnly>
                                        Name
                                    </Form.Label>
                                    <Form.Control id="inlineFormInputName" placeholder="Search ..." />
                                </Col>
                                <Col xs="auto" className="ml-auto">
                                    <Button type="submit">Search</Button>
                                </Col>
                            </Form.Row>
                        </Form>
                        <Nav className=" mb-2 ml-auto">
                            <Nav.Item>
                                <Nav.Link href="/#/signup">Sign Up</Nav.Link>
                            </Nav.Item>
                            <Nav.Item>
                                <Nav.Link href="/#/about">Login</Nav.Link>
                            </Nav.Item>
                        </Nav>
                    </Navbar.Collapse>
                </Navbar>
            </Styles>
        );
    }
}
