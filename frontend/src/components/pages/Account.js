import React, { useEffect, useState } from 'react';
import Cookies from 'js-cookie';

export function Account() {
    const [userInfo, setUserInfo] = useState(null);
    const [error, setError]       = useState(null);

    useEffect(() => {
        const userId = Cookies.get('userId');
        console.log(userId)
        if (userId) {
            fetch(`http://localhost:4000/user/${userId}`)
                .then((response) => {
                    console.log(response)
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    return response.json();
                })
                .then((data) => {
                    setUserInfo(data.user);
                })
                .catch((error) => {
                    setError(error.message);
                });
        } else {
            setError('User ID not found');
        }
    }, []);

    if (error) return <p>Error: {error}</p>;
    if (!userInfo) return <p>Loading...</p>;

    return (
        <div className="p-6">
            <h1 className="text-2xl font-bold mb-4">Account Information</h1>
            <p><b>Email</b>: {userInfo.email}</p>
            <p><b>Username</b>: {userInfo.username}</p>
        </div>
    );
}
