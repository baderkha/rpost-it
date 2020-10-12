import React, { Component } from 'react';
import { Form, Modal } from 'react-bootstrap';
import ModalFooterButtons from './Modal/ModalFooterButtons';
export default class LoginModal extends Component {
    constructor(props) {
        super();
        this.props = props;
    }
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
                    <Form>
                        <Form.Group controlId="formBasicEmail">
                            <Form.Label>Email address</Form.Label>
                            <Form.Control type="email" placeholder="Enter email" />
                        </Form.Group>

                        <Form.Group controlId="formBasicPassword">
                            <Form.Label>Password</Form.Label>
                            <Form.Control type="password" placeholder="Password" />
                        </Form.Group>

                        
                    </Form>
                </Modal.Body>
                <ModalFooterButtons submitName="Login" closeName="Close" onClose={this.props.onClose}/>
                
            </Modal>
        );
    }
}
