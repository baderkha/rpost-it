import React, { Component } from 'react';
import { Form, Modal } from 'react-bootstrap';
import { API } from '../../actions/ApiConsumer';
import ModalFooterButtons from './ModalFooterButtons';

export default class LoginModal extends Component {
    constructor(props) {
        super();
        this.props = props;
        this.api = new API('http://localhost:8080');
        this.state = {
            avatarId: '',
            password: '',
            messageErrors: '',
        };
    }
    onLoginFail = async (response) => {
        this.setState({
            messageErrors: response.Message,
        });
    };
    onFormSubmit = async (ev) => {
        ev.preventDefault();
        console.log('submit', this.state);
        this.setState({
            messageErrors: '',
        });
        let response = await this.api.login({ ...this.state });
        if (response.IsError) {
            console.log('erroring');
            this.onLoginFail(response);
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
            this.props.onLoginComplete(response.Resource);
        }
    };
    onInputchange = (event) => {
        this.setState({
            [event.target.name]: event.target.value,
        });
        console.log(this.state);
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
                    <Modal.Title id="contained-modal-title-vcenter">Login</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form onSubmit={this.onFormSubmit}>
                        <Form.Group controlId="formBasicEmail">
                            <Form.Label>Avatar Id</Form.Label>
                            <Form.Control
                                name="avatarId"
                                onChange={this.onInputchange}
                                required={true}
                                type="id"
                                placeholder="Enter Avatar Id"
                            />
                        </Form.Group>

                        <Form.Group controlId="formBasicPassword">
                            <Form.Label>Password</Form.Label>
                            <Form.Control
                                name="password"
                                type="password"
                                required={true}
                                placeholder="Password"
                                onChange={this.onInputchange}
                            />
                        </Form.Group>
                        <Form.Text style={{ color: 'red' }}>{this.state.messageErrors}</Form.Text>
                        <ModalFooterButtons
                            submitName="Login"
                            closeName="Close"
                            onClose={this.props.onClose}
                        />
                    </Form>
                </Modal.Body>
            </Modal>
        );
    }
}
