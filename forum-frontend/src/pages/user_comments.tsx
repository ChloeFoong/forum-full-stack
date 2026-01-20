import { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import "./user_posts.css";

interface Comment {
  ID: number;
  content: string;
}

export default function UserComments() {
    const { id } = useParams<{ id: string }>();
    const [comment, setComment] = useState<Comment[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const [editingCommentID, setEditingCommentID] = useState<number | null>(null);
    const [editContent, setEditContent] = useState("");

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (!token) {
        setError("You must be logged in to see your comments");
        setLoading(false);
        return;
        }

        fetch(`http://localhost:8080/users/${id}/comments`, {
        headers: { Authorization: `Bearer ${token}` },
        })
        .then((res) => {
            if (!res.ok) throw new Error("Failed to fetch comments");
            return res.json();
        })
        .then((data) => setComment(data))
        .catch((err) => setError(err.message))
        .finally(() => setLoading(false));
    }, [id]);

    const handleEditClick = (comment: Comment) => {
        setEditingCommentID(comment.ID);
        setEditContent(comment.content);
    };

    const handleUpdate = async (commentID: number) => {
        const token = localStorage.getItem("token");
        const res = await fetch(`http://localhost:8080/comments/${commentID}/update`, {
        method: "PUT",
        headers: { "Content-Type": "application/json", Authorization: `Bearer ${token}` },
        body: JSON.stringify({content: editContent,}),
        });
        if (res.ok) {
        setComment((prev) =>
            prev.map((c) =>
            c.ID === commentID ? { ...c, content: editContent} : c
            )
        );
        setEditingCommentID(null);
        } else {
        const data = await res.json();
        alert("Failed: " + (data.error || res.statusText));
        }
    };

        const handleDelete = async (commentID: number) => {
            const token = localStorage.getItem("token");
            if (!token) return alert("Not logged in");

            const res = await fetch(`http://localhost:8080/comments/${commentID}/delete`, {
                method: "DELETE",
                headers: { "Authorization": `Bearer ${token}` },
            });

            if (res.ok) {
                alert("Comment deleted!");
                setComment(comment.filter(c => c.ID !== commentID));
            } else {
                const data = await res.json();
                alert("Failed: " + (data.error || res.statusText));
            }
        };

    if (loading) return <p>Loading posts...</p>;
    if (error) return <p>{error}</p>;
    if (!comment.length) return <p>No posts yet</p>;

    return (
        <div className="user-posts-container">
        <button className="back-button" onClick={() => navigate(`/topics`)}>
            Back to topics
        </button>
        {comment.map((comment) => (
            <div key={comment.ID} className="post-card">
            {editingCommentID === comment.ID ? (
                <>
                <textarea value={editContent} onChange={(e) => setEditContent(e.target.value)} />
                <button className="edit-btn" onClick={() => handleUpdate(comment.ID)}>
                    Save
                </button>
                <button className="cancel-btn" onClick={() => setEditingCommentID(null)}>
                    Cancel
                </button>
                </>
            ) : (
                <>
                <p>{comment.content}</p>
                <button className="edit-btn" onClick={() => handleEditClick(comment)}>
                    Edit
                </button>
                <button className="delete-btn" onClick={() => handleDelete(comment.ID)}>
                    Delete
                </button>
                </>
            )}
            </div>
        ))}
        </div>
    );
}
