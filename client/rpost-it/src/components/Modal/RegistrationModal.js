import React, { Component } from 'react';
import { Form, Modal, Alert, Toast } from 'react-bootstrap';
import { API } from '../../actions/ApiConsumer';
import ModalFooterButtons from './ModalFooterButtons';

export default class RegistrationModal extends Component {
    constructor(props) {
        super();
        this.props = props;
        this.api = new API('http://localhost:8080', false);
        this.state = {
            firstName: '',
            lastName: '',
            email: '',
            dob: '',
            avatarId: '',
            password: '',
            messageErrors: '',
        };
    }
    onFormSubmit = async (ev) => {
        ev.preventDefault();
        console.log('submit');
        this.setState({
            messageErrors: '',
        });
        let response = await this.api.register({
            ...this.state,
            dob: this.state.dob + 'T00:00:00Z',
        });
        if (response.IsError) {
            console.log('erroring');
            this.onRegFail(response);
        } else {
            // clear state
            this.setState({
                firstName: '',
                lastName: '',
                email: '',
                dob: '',
                avatarId: '',
                password: '',
                messageErrors: '',
            });
            this.props.onRegistration(response.Resource);
        }
    };
    onRegFail = async (response) => {
        this.setState({
            messageErrors: response.Message,
        });
    };
    onInputchange = (event) => {
        this.setState({
            [event.target.name]: event.target.value,
        });
    };
    render() {
        return (
            <Modal
                show={this.props.isOpen}
                size="lg"
                aria-labelledby="contained-modal-title-vcenter"
                centered
            >
                <Modal.Header closeButton onClick={this.props.onClose}>
                    <Modal.Title id="contained-modal-title-vcenter">Sign Up</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form onSubmit={this.onFormSubmit}>
                        <Form.Group controlId="formBasicPII">
                            <Form.Label>First Name</Form.Label>
                            <Form.Control
                                name="firstName"
                                type="name"
                                required="true"
                                placeholder="Enter First Name"
                                onChange={this.onInputchange}
                                value={this.state.firstName}
                            />
                        </Form.Group>
                        <Form.Group controlId="formBasicPII">
                            <Form.Label>Last Name</Form.Label>
                            <Form.Control
                                name="lastName"
                                type="name"
                                required="true"
                                placeholder="Enter Last Name"
                                onChange={this.onInputchange}
                                value={this.state.lastName}
                            />
                        </Form.Group>
                        <Form.Group controlId="formBasicPII">
                            <Form.Label>Aavatar Id</Form.Label>
                            <Form.Control
                                name="avatarId"
                                type="name"
                                required="true"
                                placeholder="Enter An Avatar Id"
                                onChange={this.onInputchange}
                                value={this.state.avatarId}
                            />
                        </Form.Group>
                        <Form.Group controlId="formBasicEmail">
                            <Form.Label>Email address</Form.Label>
                            <Form.Control
                                name="email"
                                type="email"
                                required="true"
                                placeholder="Enter email"
                                onChange={this.onInputchange}
                                value={this.state.email}
                            />
                            <Form.Text className="text-muted">
                                We'll never share your email with anyone else.
                            </Form.Text>
                        </Form.Group>
                        <Form.Group controlId="dob">
                            <Form.Label>Birth Date</Form.Label>
                            <Form.Control
                                name="dob"
                                type="date"
                                placeholder="Enter Date Of Birth"
                                onChange={this.onInputchange}
                                value={this.state.dob}
                            />
                        </Form.Group>

                        <Form.Group controlId="formBasicPassword">
                            <Form.Label>Password</Form.Label>
                            <Form.Control
                                name="password"
                                type="password"
                                required="true"
                                placeholder="Password"
                                onChange={this.onInputchange}
                                value={this.state.password}
                            />
                        </Form.Group>
                        <Form.Group controlId="formBasicCheckbox">
                            <Form.Check
                                type="checkbox"
                                label="I Agree to The Terms and Conditions"
                            />
                        </Form.Group>
                        <Form.Text style={{ color: 'red' }}>{this.state.messageErrors}</Form.Text>
                        <ModalFooterButtons
                            submitName="SignUp"
                            closeName="Close"
                            onClose={this.props.onClose}
                        />
                    </Form>
                </Modal.Body>
            </Modal>
        );
    }
}
