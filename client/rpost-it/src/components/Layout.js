import Container from 'react-bootstrap/Container';

import React, { Component } from 'react';
import Styled from 'styled-components';
const Style = Styled.div`
    .container{
        padding:20px;
    }
`;
export default class Layout extends Component {
    constructor(props) {
        super();
        this.props = props;
    }
    render() {
        return <Style><Container>{this.props.children}</Container></Style>;
    }
}
