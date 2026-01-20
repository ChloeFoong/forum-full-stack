import { useParams, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import "./create_comment.css";

export default function CommentForm() {
  const { postId } = useParams<{ postId: string }>();
  const numericPostId = Number(postId);

  const [content, setContent] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();


  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!content.trim()) {
      setMessage("Comment cannot be empty");
      return;
    }

    try {
      const token = localStorage.getItem("token");
      if (!token) {
        setMessage("You must be logged in to comment");
        return;
      }

      const res = await fetch(
        `http://localhost:8080/posts/${numericPostId}/comments`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
          },
          body: JSON.stringify({ content, post_id: numericPostId }),
        }
      );

      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || res.statusText);
      }

      setMessage("Comment created successfully!");
      setContent("");
    } catch (err: any) {
      setMessage(err.message || "Failed to create comment");
    }
  };


  return (
    <div className="comment-container">
      <h3 style={{ textAlign: "center" }}>Write a Comment</h3>
      <form onSubmit={handleSubmit}>
        <textarea
          className="content-area"
          placeholder="Type your comment here..."
          value={content}
          onChange={(e) => setContent(e.target.value)}
        />
        <button type="submit">
          Submit Comment
        </button>
      </form>
      <button onClick={() => navigate(`/topics`)} style={{padding:5, marginTop: 10, display:"block"}}>Back to topics</button>

      {message && <p className="comment-message">{message}</p>}
    </div>
  );
}
