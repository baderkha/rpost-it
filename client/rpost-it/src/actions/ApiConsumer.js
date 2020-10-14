import axios from 'axios';
export class API {
    constructor(url, logging = false) {
        this.url = url;
        this.logging = logging;
    }
    /**
     * Will either return an object or string , if string  == error
     * @param {*} param0
     */
    register({ firstName, lastName, email, dob, avatarId, password }) {
        return axios
            .post(
                `${this.url}/api/accounts`,
                {
                    firstName,
                    lastName,
                    email,
                    dob,
                    avatarId,
                    password,
                },
                {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                }
            )
            .then(({ data }) => {
                return data;
            })
            .catch((err) => {
                console.log(err);
                return err.response.data;
            });
    }
    login({ avatarId, password }) {
        return axios
            .post(
                `${this.url}/api/accounts/jwt`,
                { avatarId, password },
                {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                }
            )
            .then(({ data }) => {
                return data;
            })
            .catch((err) => {
                return err.response.data;
            });
    }
    getPostById(postId) {
        return axios
            .get(`${this.url}/api/posts/${postId}`)
            .then(({ data }) => {
                return data.Resource;
            })
            .catch((err) => {
                console.log(err);
                return err;
            });
    }
    getPostsForCommunityId(commnityId) {
        return axios
            .get(`${this.url}/api/community/${commnityId}/posts`)
            .then(({ data }) => {
                return data.Resource;
            })
            .catch((err) => {
                console.log(err);
                return err;
            });
    }

    getCommunityById(communityId) {
        return axios
            .get(`${this.url}/api/community/${communityId}`)
            .then(({ data }) => {
                return data.Resource;
            })
            .catch((err) => {
                console.log(err);
                return err;
            });
    }

    getAccountInfoInternal(jwt) {
        return axios
            .get(`${this.url}/api/jwt/accounts`, {
                headers: {
                    Authorization: `Bearer ${jwt}F`,
                },
            })
            .then(({ data }) => {
                return data.Resource;
            })
            .catch((err) => {
                console.log(err);
                return err;
            });
    }
}
