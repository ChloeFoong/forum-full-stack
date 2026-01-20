import React, { useState, useEffect } from "react";
import "./create_post.css";
import { useNavigate } from "react-router-dom";

interface Topic {
  ID: number;
  name: string;
}

export default function CreatePost() {
    const [heading, setHeading] = useState("");
    const [content, setContent] = useState("");
    const [tag, setTag] = useState("");
    const [topicId, setTopicId] = useState<number | "">("");
    const [topics, setTopics] = useState<Topic[]>([]);
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        const fetchTopics = async () => {
        try {
            const res = await fetch("http://localhost:8080/topics");
            const data: Topic[] = await res.json();
            setTopics(data);
        } catch (err) {
            setMessage("Error fetching topics from backend");
        }
        };
        fetchTopics();
    }, []);

    const handleSubmit = async (e: React.FormEvent) => {
        const token = localStorage.getItem("token");
        e.preventDefault();
        if (!topicId) {
        setMessage("Please select a topic");
        return;
        }
        try {
        const res = await fetch("http://localhost:8080/posts", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`},
            body: JSON.stringify({
            heading,
            content,
            tag: tag.split(",").map((t) => ({ Name: t.trim() })),
            topic_id: Number(topicId),
            }),
        });
        const data = await res.json();
        if (res.ok) {
            setMessage("Post created successfully!");
            setHeading("");
            setContent("");
            setTag("");
        } else {
            setMessage("Failed: " + (data.error || res.statusText));
        }
        } catch {
            setMessage("Error connecting to backend");
        }
    };

    return (
        <div className="create-post-container">
        <h2 className="create-post-title">Create a New Post</h2>
        <form onSubmit={handleSubmit} className="create-post-form">
            <input
            type="text"
            placeholder="Heading"
            value={heading}
            onChange={(e) => setHeading(e.target.value)}
            className="create-post-input"
            required
            />
            <textarea
            placeholder="Content"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            className="create-post-textarea"
            required
            />
            <input
            type="text"
            placeholder="Tag (optional)"
            value={tag}
            onChange={(e) => setTag(e.target.value)}
            className="create-post-input"
            />
            <select
                value={topicId}
                onChange={(e) => setTopicId(Number(e.target.value))}
                required
                >
                <option value="">Select a topic</option>

                {topics.map((t) => (
                    <option key={t.ID} value={t.ID}>
                    {t.name}
                    </option>
                ))}
                </select>
            <button type="submit" className="create-post-button">
            Create Post
            </button>
        </form>
        {message && <p className="create-post-message">{message}</p>}

        <button onClick={() => navigate(`/topics/${topicId}/posts`)} 
            style={
                {padding: 10, 
                margin: "20px auto", 
                display: "block",
                fontSize: "16px",
                borderRadius: "5px"}}>See post</button>
        </div>
    );
}
