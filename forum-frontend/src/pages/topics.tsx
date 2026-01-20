import { useState, ChangeEvent, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "./topics.css";
import {jwtDecode} from "jwt-decode";

interface Topic {
    ID: number;
    name: string;
}


interface TokenPayload {
    username: string;
    id: number;
    exp: number;
  }

export default function MainPage () {
    const [topics, setTopics] = useState<Topic[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");
    const navigate = useNavigate();
    
    const token = localStorage.getItem("token");
    let userID: number | null = null;
    
    if (token) {
      try {
        const decoded = jwtDecode<TokenPayload>(token);
        userID = decoded.id;
      } catch (err) {
        console.error("Invalid token", err);
      }
    }
    useEffect(() => {

        fetch("http://localhost:8080/topics") 
          .then((res) => {
            if (!res.ok) throw new Error("Failed to fetch topics");
            return res.json();
          })
          .then((data: Topic[]) => {
            setTopics(data);
            setLoading(false);
          })
          .catch((err) => {
            setError(err.message);
            setLoading(false);
          });
      }, []);
    if (loading) {
        return <p>Loading topics...</p>;
    }
    if (error) {
        return <p>Error: {error}</p>;
    }
    return (
        <div>
        <div>
          <ul>
            <li><button onClick={() => navigate(`/topics`)}>Home</button></li>
            <li><button onClick={() => navigate(`/posts`)}>Create a post</button></li>
            <li><button onClick={() => navigate(`/`)}>Logout</button></li>
            <li><button onClick={() => navigate(`/users/${userID}/posts`)}>My posts</button></li>
            <li><button onClick={() => navigate(`/users/${userID}/comments`)}>My comments</button></li>
          </ul>
          <h1 className="mainTitle">Topics</h1>
        </div>
      
          <ul className="topic-grid">
            {topics.map((topic) => (
              <li key={topic.ID} className="topic-card">
                <h3 className="topic-title">{topic.name}</h3>
      
                <div className="topic-actions">
                  <button
                    onClick={() => navigate(`/topics/${topic.ID}/posts`)}
                  >
                    Posts
                  </button>

                </div>
              </li>
            ))}
          </ul>
        </div>
      );
      
      
}