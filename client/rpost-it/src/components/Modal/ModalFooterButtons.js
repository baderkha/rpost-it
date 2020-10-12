import React, { Component } from 'react';
import { Button, Modal } from 'react-bootstrap';
import styled from 'styled-components';
const Styles = styled.div`
    .signup-footer {
        width: 100%;
    }
    .cls-btn{
        float:right;
    }
`;
export default class ModalFooterButtons extends Component {
    constructor(props) {
        super();
        this.props = props;
    }
    render() {
        return (
            <Styles>
                <Modal.Footer>
                    <div className="signup-footer">
                        <Button variant="primary" type="submit" onClick={this.props.onSubmit}>
                            {this.props.submitName}
                        </Button>
                        <div className="cls-btn">
                            <Button onClick={this.props.onClose}>{this.props.closeName}</Button>
                        </div>
                    </div>
                </Modal.Footer>
            </Styles>
        );
    }
}
