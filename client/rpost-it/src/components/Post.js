import React, { Component } from 'react';
import {Card, Form, FormControl, FormGroup } from 'react-bootstrap';
import PostButton from './Buttons/PostButton';
import Styled  from 'styled-components';

const Style = Styled.div`
    .card{
        margin-bottom:20px;
    }
    .form-control{
        height:100px;
    }
    .inline-post{
        display:wrap;
        text-align:right;
    }
    .btn{
        margin-top : 10px;   
        float: right;
    }
`;

export default class Post extends Component {
    render() {
        return (
            <Style>
                <Card>
                    <Card.Header>Post Something</Card.Header>
                    <Card.Body>
                        <Form>
                            <FormGroup>
                                <div className="inline-post">
                                    <FormControl
                                        componentClass="textarea"
                                        placeholder="Post Something"
                                    />
                                </div>
                            </FormGroup>
                            <PostButton label="Post" onClick={()=>alert('post')}/>
                        </Form>
                    </Card.Body>
                </Card>
            </Style>
        );
    }
}
