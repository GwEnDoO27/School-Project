// Function to get a cookie by name
export function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) {
        return parts.pop().split(';').shift(); // Return the actual cookie value
    }
    return null; // Return null if the cookie is not found
}

// Get the session cookie value
export let username = "";

export function DecodeUsername(sessionCookie) {
    if (sessionCookie) {
        // Extract the username from the cookie value
        // Use '%3A' to split the session-token
        const parts = sessionCookie.split('%3A');
        if (parts.length === 2) {

            // Decode username from cookie
            username = decodeURIComponent(parts[1]);
        }
    }
    return username
}

// Function to delete a cookie by setting its expiration date to a past date
export function deleteCookie(name) {
    document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
}

// Function to notify the backend to delete the session
export async function deleteSession() {
    const sessionCookie = getCookie('session-token');
    if (sessionCookie) {
        await fetch('/delete-session', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ sessionToken: sessionCookie })
        });
    }
}
