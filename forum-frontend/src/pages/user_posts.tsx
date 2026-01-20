import { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import "./user_posts.css";

interface Post {
  ID: number;
  heading: string;
  content: string;
  tag: string;
}

export default function UserPosts() {
  const { id } = useParams<{ id: string }>();
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const [editingPostID, setEditingPostID] = useState<number | null>(null);
  const [editHeading, setEditHeading] = useState("");
  const [editContent, setEditContent] = useState("");
  const [editTag, setEditTag] = useState("");

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setError("You must be logged in to see your posts");
      setLoading(false);
      return;
    }

    fetch(`http://localhost:8080/users/${id}/posts`, {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch posts");
        return res.json();
      })
      .then((data) => setPosts(data))
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false));
  }, [id]);

  const handleEditClick = (post: Post) => {
    setEditingPostID(post.ID);
    setEditHeading(post.heading);
    setEditContent(post.content);
    setEditTag(post.tag);
  };

  const handleUpdate = async (postID: number) => {
    const token = localStorage.getItem("token");
    const res = await fetch(`http://localhost:8080/posts/${postID}/update`, {
      method: "PUT",
      headers: { "Content-Type": "application/json", Authorization: `Bearer ${token}` },
      body: JSON.stringify({ heading: editHeading, content: editContent, tag: editTag }),
    });
    if (res.ok) {
      setPosts((prev) =>
        prev.map((p) =>
          p.ID === postID ? { ...p, heading: editHeading, content: editContent, tag: editTag } : p
        )
      );
      setEditingPostID(null);
    } else {
      const data = await res.json();
      alert("Failed: " + (data.error || res.statusText));
    }
  };

    const handleDelete = async (postID: number) => {
        const token = localStorage.getItem("token");
        if (!token) return alert("Not logged in");

        const res = await fetch(`http://localhost:8080/posts/${postID}/delete`, {
            method: "DELETE",
            headers: { "Authorization": `Bearer ${token}` },
        });

        if (res.ok) {
            alert("Post deleted!");
            setPosts(posts.filter(p => p.ID !== postID));
        } else {
            const data = await res.json();
            alert("Failed: " + (data.error || res.statusText));
        }
    };

  if (loading) return <p>Loading posts...</p>;
  if (error) return <p>{error}</p>;
  if (!posts.length) return <p>No posts yet</p>;

  return (
    <div className="user-posts-container">
      <button className="back-button" onClick={() => navigate(`/topics`)}>
        Back to topics
      </button>
      <h2>My Posts</h2>
      {posts.map((post) => (
        <div key={post.ID} className="post-card">
          {editingPostID === post.ID ? (
            <>
              <input value={editHeading} onChange={(e) => setEditHeading(e.target.value)} />
              <textarea value={editContent} onChange={(e) => setEditContent(e.target.value)} />
              <input value={editTag} onChange={(e) => setEditTag(e.target.value)} />
              <button className="edit-btn" onClick={() => handleUpdate(post.ID)}>
                Save
              </button>
              <button className="cancel-btn" onClick={() => setEditingPostID(null)}>
                Cancel
              </button>
            </>
          ) : (
            <>
              <h3>{post.heading}</h3>
              <p>{post.content}</p>
              {post.tag && <p>#{post.tag}</p>}
              <button className="edit-btn" onClick={() => handleEditClick(post)}>
                Edit
              </button>
              <button className="delete-btn" onClick={() => handleDelete(post.ID)}>
                Delete
              </button>
            </>
          )}
        </div>
      ))}
    </div>
  );
}
