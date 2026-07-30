class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<int> st;

        for (const string& tok : tokens) {
            if (tok == "+" || tok == "-" || tok == "*" || tok == "/") {
                int b = st.top(); st.pop();
                int a = st.top(); st.pop();
                if (tok == "+") st.push(a + b);
                else if (tok == "-") st.push(a - b);
                else if (tok == "*") st.push(a * b);
                else st.push(a / b); // C++ integer division truncates toward zero
            } else {
                st.push(stoi(tok));
            }
        }

        return st.top();
    }
};