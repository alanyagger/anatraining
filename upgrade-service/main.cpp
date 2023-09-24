#include <bits/stdc++.h>

using namespace std;

// Uncomment the following line if you need it
// typedef long long ll;

int main() {
	unsigned long long n, m, k,temp,sig,minsum=0;
	long long did=0;
	cin.sync_with_stdio(false);
	queue<unsigned long long> q1;
	queue<unsigned long long> q2;
	cin >> n >> m >> k;
	for (unsigned long long i = 1; i <= n; i++) {
		unsigned long long ai;
		cin >> ai;
		q1.push(ai);    
	}
	for (unsigned long long j=1;j<=m;j++)
	{
		minsum+=q1.front();
		q2.push(q1.front());
		q1.pop();
	}
	temp=1;
	sig=1;
	while(!q1.empty())
	{
		temp++;
		did=did+q1.front()-q2.front();
		q2.push(q1.front());
		q1.pop();
		q2.pop();
		if (did<0)
		{
			sig=temp;
			minsum+=did;
			did=0;
		}
	}
	if ( minsum< m*k)
	{
		cout << sig<<' '<<sig+m-1<<endl;
	}
    else{
    	cout<<0<<' '<<0<<endl;
	}
	return 0;
}
