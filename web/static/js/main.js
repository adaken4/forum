// async function toggleReaction(postID, type) {
//     try {
//         const response = await fetch(`/like?post_id=${postID}&type=${type}`, {
//             method: "POST",
//             credentials: "include"
//         });

//         if (!response.ok) {
//             const errorText = await response.text();
//             alert("Error: " + errorText);
//             return;
//         }

//         const data = await response.json();
//         document.getElementById(`likes-${postID}`).innerText = data.likes;
//         document.getElementById(`dislikes-${postID}`).innerText = data.dislikes;
//     } catch (error) {
//         console.error("Error toggling like:", error);
//     }
// }

// Function to handle like/dislike for a post
async function reactToPost(userId, postId, likeType) {
    try {
        const response = await fetch("/like", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ 
                user_id: userId,
                post_id: postId,
                like_type: likeType
            })
        });

        const result = await response.json();
        console.log(result.message);

        if (response.ok) {
            updateReactionUI(postId, null, result.likes, result.dislikes, result.userReaction);  // Update UI
        }
    } catch (error) {
        console.error("Error:", error);
    }
}

// Function to handle like/dislike for a comment
async function reactToComment(userId, commentId, likeType) {
    try {
        const response = await fetch("/like", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ 
                user_id: userId,
                comment_id: commentId,
                like_type: likeType
            })
        });

        const result = await response.json();
        console.log(result.message);

        if (response.ok) {
            updateReactionUI(null, commentId, result.likes, result.dislikes, result.userReaction);  // Update UI
        }
    } catch (error) {
        console.error("Error:", error);
    }
}

// Function to update the UI instantly without reloading
function updateReactionUI(postId, commentId, likes, dislikes, userReaction) {
    let likeBtn, dislikeBtn, likeCount, dislikeCount;

    if (postId) {
        likeBtn = document.querySelector(`#like-post-${postId}`);
        dislikeBtn = document.querySelector(`#dislike-post-${postId}`);
        likeCount = document.querySelector(`#like-count-${postId}`);
        dislikeCount = document.querySelector(`#dislike-count-${postId}`);
    } else if (commentId) {
        likeBtn = document.querySelector(`#like-comment-${commentId}`);
        dislikeBtn = document.querySelector(`#dislike-comment-${commentId}`);
        likeCount = document.querySelector(`#like-count-${commentId}`);
        dislikeCount = document.querySelector(`#dislike-count-${commentId}`);
    }

    if (!likeBtn || !dislikeBtn || !likeCount || !dislikeCount) {
        console.error("UI elements not found");
        return;
    }

    // if (likeType === "like") {
    //     if (likeBtn.classList.contains("active")) {
    //         likeBtn.classList.remove("active");
    //         likeCount.innerText = parseInt(likeCount.innerText) - 1;
    //     } else {
    //         likeBtn.classList.add("active");
    //         likeCount.innerText = parseInt(likeCount.innerText) + 1;
    //         if (dislikeBtn.classList.contains("active")) {
    //             dislikeBtn.classList.remove("active");
    //             dislikeCount.innerText = parseInt(dislikeCount.innerText) - 1;
    //         }
    //     }
    // } else if (likeType === "dislike") {
    //     if (dislikeBtn.classList.contains("active")) {
    //         dislikeBtn.classList.remove("active");
    //         dislikeCount.innerText = parseInt(dislikeCount.innerText) - 1;
    //     } else {
    //         dislikeBtn.classList.add("active");
    //         dislikeCount.innerText = parseInt(dislikeCount.innerText) + 1;
    //         if (likeBtn.classList.contains("active")) {
    //             likeBtn.classList.remove("active");
    //             likeCount.innerText = parseInt(likeCount.innerText) - 1;
    //         }
    //     }
    // }
    // Update counts with real backend data
    likeCount.innerText = likes;
    dislikeCount.innerText = dislikes;

    // Reset button states
    likeBtn.classList.remove("active");
    dislikeBtn.classList.remove("active");

    if (userReaction === "like") {
        likeBtn.classList.add("active");
    } else if (userReaction === "dislike") {
        dislikeBtn.classList.add("active");
    }
}
