import React, { Component } from 'react';
import { Nav, Navbar, Form, Col } from 'react-bootstrap';
import styled from 'styled-components';
import Registration from './Modal/RegistrationModal';
import LoginModal from './Modal/LoginModal';
import SearchButton from './Buttons/SearchButton';
const Styles = styled.div`
    .navbar {
        background-color: #222;
    }
    .navbar-brand {
        color: white;
        margin-left: 5px;
        &:hover {
            color: white;
            opacity: 0.5;
        }
        &:visited {
            color: white;
            opactiy: 1;
        }
        &:focus {
            color: white;
            opactiy: 1;ƒ
        }
        &:unfocus {
            color: white;
            opactiy: 1;
        }
    }
    .navbar-nav .nav-link {
        color: white;
        &:hover {
            color: white;
            opacity: 0.5;
        }
        &:visited {
            color: white;
            opactiy: 1;
        }
        &:focus {
            color: white;
            opactiy: 1;
        }
        &:unfocus {
            color: white;
            opactiy: 1;
        }
    }

    .navbar-collapse {
        padding: 5px;
    }
    .search-bar-form {
        flex-grow: 1;
        width: 100%;
        max-width: 690px;
        margin: auto;
    }
    .form-control {
        width: 100%;
    }
    .nav-item {
        margin-top: 10px;
    }
`;

export default class NavigationBar extends Component {
    constructor(props) {
        super();
        this.props = props;
        this.state = {
            isRegistrationModalOpen: false,
            isLoginModalOpen: false,
            isLoggedIn: true,
        };
    }
    openRegModal = () => this.setState({ isRegistrationModalOpen: true });
    closeRegModal = () => this.setState({ isRegistrationModalOpen: false });
    openLogModal = () => this.setState({ isLoginModalOpen: true });
    closeLogModal = () => this.setState({ isLoginModalOpen: false });
    showLoginState = () => this.setState({ isLoggedIn: true });
    hideLoginState = () => this.setState({ isLoggedIn: false });
    showLoginStateHideModal = () => this.setState({ isLoginModalOpen: false, isLoggedIn: true });

    render() {
        return (
            <Styles>
                <Navbar expand="lg">
                    <Navbar.Brand className="mb-2" href="/">
                        rpost-it
                    </Navbar.Brand>
                    <Navbar.Toggle area-aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Form className="search-bar-form">
                            <Form.Row>
                                <Col>
                                    <Form.Control
                                        id="inlineFormInputName"
                                        placeholder="Search ..."
                                    />
                                </Col>
                                <Col xs="auto">
                                    <SearchButton onClick={() => alert('clicked')} />
                                </Col>
                            </Form.Row>
                        </Form>
                        <Nav>
                            <Nav.Item hidden={this.state.isLoggedIn}>
                                <Nav.Link onClick={this.openRegModal}>Sign Up</Nav.Link>
                            </Nav.Item>
                            <Nav.Item hidden={this.state.isLoggedIn}>
                                <Nav.Link onClick={this.openLogModal}>Login</Nav.Link>
                            </Nav.Item>

                            <Nav.Item hidden={!this.state.isLoggedIn}>
                                <Nav.Link>Settings</Nav.Link>
                            </Nav.Item>
                            <Nav.Item hidden={!this.state.isLoggedIn}>
                                <Nav.Link>Ahmad Baderkhan</Nav.Link>
                            </Nav.Item>
                            <Nav.Item hidden={!this.state.isLoggedIn}>
                                <Nav.Link onClick={this.hideLoginState}>Log out</Nav.Link>
                            </Nav.Item>
                        </Nav>
                    </Navbar.Collapse>
                </Navbar>
                <Registration
                    isOpen={this.state.isRegistrationModalOpen}
                    onClose={this.closeRegModal}
                />
                <LoginModal
                    isOpen={this.state.isLoginModalOpen}
                    onLogin={this.showLoginStateHideModal}
                    onClose={this.closeLogModal}
                />
            </Styles>
        );
    }
}
