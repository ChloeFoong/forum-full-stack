import { Routes, Route } from "react-router-dom";
import Login from "./pages/login";
import Create from "./pages/create";
import MainPage from "./pages/topics";
import TopicPosts from "./pages/topic_posts";
import CreatePost from "./pages/create_post";
import UserPosts from "./pages/user_posts";
import CommentForm from "./pages/create_comment";
import PostComments from "./pages/view_comments";
import UserComments from "./pages/user_comments";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="/users" element={<Create />} />
      <Route path="/topics" element={<MainPage />} />
      <Route path="/topics/:topicId/posts" element={<TopicPosts key={window.location.pathname} />} />
      <Route path="/posts" element={<CreatePost />} />
      <Route path="/users/:id/posts" element={<UserPosts />} />
      <Route path="/posts/:postId/comments" element={<CommentForm />} />
      <Route path="/posts/:postId/allcomment" element={<PostComments />} />
      <Route path="/users/:id/comments" element={<UserComments />} />
    </Routes>
  );
}

export default App;
