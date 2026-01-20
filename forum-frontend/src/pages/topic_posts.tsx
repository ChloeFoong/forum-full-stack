import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";

interface Topic {
  ID: number;
  name: string;
}

interface Post {
  ID: number;
  heading: string;
  content: string;
  tag: string;
}

export default function TopicPosts() {
    const { topicId } = useParams<{ topicId: string }>();
    const [topic, setTopic] = useState<Topic | null>(null);
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        if (!topicId) return;

    setLoading(true);
    setTopic(null);
    setPosts([]);
  
    const fetchTopicAndPosts = async () => {
        try {
            const topicRes = await fetch(`http://localhost:8080/topics/${topicId}`);
            if (!topicRes.ok) throw new Error("Topic not found");
            const topicData: Topic = await topicRes.json();
            setTopic(topicData);
    
            const postsRes = await fetch(`http://localhost:8080/topics/${topicId}/posts`);
            if (!postsRes.ok) throw new Error("Posts not found");
            const postsData: Post[] = await postsRes.json();
            setPosts(postsData);
        } catch (err) {
            console.error(err);
        } finally {
            setLoading(false);
        }
    };
  
    fetchTopicAndPosts();
    }, [topicId]);
  

    if (loading) return <p>Loading posts...</p>;
    if (!topic) return <p>Topic not found</p>;

    return (
        <div>
            
            <h2>{topic.name}</h2>

            <button onClick={() => navigate(`/topics`)}>Back to topics</button>
            
            {posts.length === 0 ? (
                <p>No posts yet</p>
            ) : (
                posts.map((p) => (
                <div key={p.ID} style={{border: "2px solid gray", margin: "10px auto", display: "block",padding:10}}>
                    <h3 style={{marginLeft: 20, display: "block"}}>{p.heading}</h3>
                    <p style={{marginLeft: 20, display: "block"}}>{p.content}</p>
                    <button onClick={() => navigate(`/posts/${p.ID}/allcomment`)}>View Comments</button>
                    <button onClick={() => navigate(`/posts/${p.ID}/comments`)}>Comment</button>
                </div>
                ))
            )}
        </div>
    );
}
