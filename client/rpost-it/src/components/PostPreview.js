import React, { Component } from 'react';
import { Card } from 'react-bootstrap';
import Styled from 'styled-components';
import LikeButton from './Buttons/LikeButton';
import DisLikeButton from './Buttons/DisLikeButton';
import CommentButton from './Buttons/CommentButton';
const Style = Styled.div`
    .card{
        margin-bottom:20px;
    }
    .cart-title{
        text-align:center;
    }
    .btn-lk{
        float : left;
        margin-right : 10px;
    }
    .btn-lk button{
        margin-right:10px;
    }
    .btn-cmnt{
        float : right;
    }
    .left-text-header{
        float : left;
    }
    .right-text-header{
        float : ;ef;
    }
   
`;

export default class PostPreview extends Component {
    constructor(props) {
        super();
        this.props = props;
    }
    render() {
        return (
            <Style>
                <Card>
                    <Card.Header>
                        <p className="left-text-header">Posted By : {this.props.postedBy}</p>
                    </Card.Header>
                    <Card.Body>
                        <Card.Title as="h5">{this.props.postTitle}</Card.Title>
                    </Card.Body>
                    <Card.Footer>
                        <div className="btn-lk">
                            <LikeButton counter="10" />
                            <DisLikeButton counter="20" />
                        </div>
                        <div className="btn-cmnt">
                            <CommentButton counter="20" />
                        </div>
                    </Card.Footer>
                </Card>
            </Style>
        );
    }
}
