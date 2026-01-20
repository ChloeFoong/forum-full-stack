import { useParams, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

interface Comment {
    ID: number;
    user_id: number;
    post_id: number;
    content: string;
}

export default function PostComments() {
    const { postId } = useParams<{ postId: string }>();
    const [comments, setComments] = useState<Comment[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        fetch(`http://localhost:8080/posts/${postId}/allcomment`)
            .then(res => {
                if (!res.ok) throw new Error("Failed to fetch comments");
                return res.json();
            })
            .then(data => setComments(data))
            .catch(err => setError(err.message))
            .finally(() => setLoading(false));
    }, [postId]);

    if (loading) return <p>Loading comments...</p>;
    if (error) return <p>Error: {error}</p>;
    if (comments.length === 0) return <p>No comments yet.</p>;

    return (
        <div>
            <h3>Comments</h3>
            {comments.map(c => (
                <div key={c.ID} style={{ border: "1px solid gray", padding: "5px", margin: "5px 0" }}>
                    <p>{c.content}</p>
                </div>
            ))}
        </div>
    );
}
