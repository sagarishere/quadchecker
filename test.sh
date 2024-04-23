#! /bin/bash

echo """Test 1
Command:
./quadA 3 3 | ./quadchecker
Expected output:
[quadA] [3] [3]
"""
./quadA 3 3 | ./quadchecker

echo """Test 2
Command:
./quadB 3 3 | ./quadchecker
Expected output:
[quadA] [3] [3]
"""
./quadB 3 3 | ./quadchecker

echo """Test 3
Command:
./quadC 1 1 | ./quadchecker
Expected output:
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
"""
./quadC 1 1 | ./quadchecker

echo """Test 4
Command:
./quadE 1 2 | ./quadchecker
Expected output:
[quadC] [1] [2] || [quadE] [1] [2]
"""
./quadE 1 2 | ./quadchecker

echo """Test 5
Command:
./quadE 2 1 | ./quadchecker
Expected output:
[quadD] [2] [1] || [quadE] [2] [1]
"""
./quadE 2 1 | ./quadchecker

echo """Test 6
Command:
./quadC 2 1 | ./quadchecker
Expected output:
[quadC] [2] [1]
"""
./quadC 2 1 | ./quadchecker

echo """Test 7
Command:
./quadD 1 2 | ./quadchecker
Expected output:
[quadD] [1] [2]
"""
./quadD 1 2 | ./quadchecker

echo """Test 8
Command:
./quadE 1 1 | ./quadchecker
Expected output:
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
"""
./quadE 1 1 | ./quadchecker

echo """Test 9
Command:
echo 0 0 | ./quadchecker
Expected output:
Not a quad function
"""
echo 0 0 | ./quadchecker

echo """Test 10
Command:
echo -n "o--" | ./quadchecker
Expected output:
Not a quad function
"""
echo -n "o--" | ./quadchecker

echo """Test 11
Command:
echo -n "/****" | ./quadchecker
Expected output:
Not a quad function
"""
echo -n "/****" | ./quadchecker

echo """Test 12
Command:
echo -n "ABBB" | ./quadchecker
Expected output:
Not a quad function
"""
echo -n "ABBB" | ./quadchecker

echo """Test 13
Command:
echo -n "ABBBA"$'\n'"B"$'\n'"B" | ./quadchecker
Expected output:
Not a quad function
"""
echo -n "ABBBA"$'\n'"B"$'\n'"B" | ./quadchecker

cat <<'EOF'
Test 14
Command:
echo -n "o--o"$'\n'"|"'\n'"o" | ./quadchecker
Expected output:
Not a quad function
EOF
echo
echo -n "o--o"$'\n'"|"$'\n'"o" | ./quadchecker

echo """All tests complete!"""